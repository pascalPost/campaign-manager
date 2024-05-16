import { ReactElement, useEffect, useState } from "react";
import { ChevronDown, ChevronRight } from "lucide-react";
import { useQuery } from "react-query";
import { client } from "@/lib/api/client.ts";

type File = {
  type: "file";
  path: string;
};

type Folder = {
  type: "folder";
  path: string;
  childPaths?: string[];
  isFolded: boolean;
};

type SelectedFileProps = {
  selectedFile?: string;
  handleChangeSelectedFile: (fileId: string) => void;
};

function FileTreeFile({
  file,
  selectedFileProps,
}: {
  file: File;
  selectedFileProps: SelectedFileProps;
}): ReactElement {
  const fileName = file.path.split("/").pop() || file.path;

  return (
    <li className="flex flex-row">
      <div className="w-5" />
      <div
        onClick={() => selectedFileProps.handleChangeSelectedFile(file.path)}
        className="hover:cursor-pointer"
      >
        {fileName}
      </div>
    </li>
  );
}

async function callGetFileTree(path: string, signal?: AbortSignal) {
  if (path === "/") {
    return client.GET("/fileTree", {
      signal,
    });
  } else {
    const pathWithoutLeadingSlash = path.replace(/^\/+/, "");
    return client.GET("/fileTree/{path}", {
      params: {
        path: {
          path: pathWithoutLeadingSlash,
        },
      },
      signal,
    });
  }
}

async function getFileTree(path: string, signal?: AbortSignal) {
  const { data, error } = await callGetFileTree(path, signal);

  if (error) throw error;

  if (!data) {
    return [];
  }

  return data.map((entry): File | Folder => {
    if (entry.isDir) {
      return {
        type: "folder",
        path: entry.name,
        isFolded: true,
      };
    } else {
      return {
        type: "file",
        path: entry.name,
      };
    }
  });
}

function FileTreeFolder({
  path,
  tree,
  onChangeFold,
  onUpdateFolder,
  selectedFileProps,
}: {
  path: string;
  tree: Map<string, Folder | File>;
  onChangeFold: (path: string, state: boolean) => void;
  onUpdateFolder: (path: string, data: (File | Folder)[]) => void;
  selectedFileProps: SelectedFileProps;
}) {
  const folder = tree.get(path) as Folder;
  const isFolded = folder.isFolded;

  const query = useQuery({
    queryKey: ["getFileTree", folder.path],
    queryFn: ({ signal }) => getFileTree(folder.path, signal),
    enabled: !isFolded && folder.childPaths == undefined,
    staleTime: Infinity,
    cacheTime: Infinity,
  });

  useEffect(() => {
    if (query.data) {
      onUpdateFolder(path, query.data);
    }
  }, [query.data]);

  const folderName = folder.path.split("/").pop() || folder.path;

  if (isFolded) {
    return (
      <>
        <li className="flex flex-row items-center">
          <ChevronRight
            className="mt-1 h-4 w-5 hover:cursor-pointer"
            onClick={() => onChangeFold(path, false)}
          />
          <div>{folderName}</div>
        </li>
      </>
    );
  }

  if (query.isLoading) {
    return "Loading...";
  }

  if (query.isError) {
    return "Error.";
  }

  if (query.isSuccess) {
    const folder = tree.get(path) as Folder;

    return (
      <>
        <li className="flex flex-row items-center">
          <ChevronDown
            className="mt-1 h-4 w-5 hover:cursor-pointer"
            onClick={() => onChangeFold(path, true)}
          />
          {folderName}
        </li>
        <ul className="px-4">
          {folder.childPaths?.map((entry) => {
            const res = tree.get(entry);
            if (!res) return;
            if (res.type === "file") {
              return (
                <FileTreeFile
                  key={entry}
                  file={res}
                  selectedFileProps={selectedFileProps}
                />
              );
            }
            if (res.type === "folder") {
              return (
                <FileTreeFolder
                  key={entry}
                  path={entry}
                  tree={tree}
                  onChangeFold={onChangeFold}
                  onUpdateFolder={onUpdateFolder}
                  selectedFileProps={selectedFileProps}
                />
              );
            }
          })}
        </ul>
      </>
    );
  }
}

function FileTree({
  selectedFileProps,
}: {
  selectedFileProps: SelectedFileProps;
}) {
  const [tree, setTree] = useState(
    new Map<string, Folder | File>([
      [
        "/",
        { type: "folder", path: "/", childPaths: undefined, isFolded: true },
      ],
    ]),
  );

  function handleChangeFold(path: string, state: boolean) {
    const folder = tree.get(path) as Folder;
    folder.isFolded = state;
    setTree(new Map(tree.set(path, folder)));
  }

  function handleUpdateFolder(path: string, data: (File | Folder)[]) {
    const root = path === "/" ? path : path + "/";
    const folder = tree.get(path) as Folder;

    data.forEach((e) => {
      e.path = root + e.path;
      tree.set(e.path, e);
    });

    folder.childPaths = data.map((e) => e.path);

    setTree(new Map(tree));
  }

  return (
    <>
      <ul className="px-1">
        <FileTreeFolder
          key={"/"}
          path={"/"}
          tree={tree}
          onChangeFold={handleChangeFold}
          onUpdateFolder={handleUpdateFolder}
          selectedFileProps={selectedFileProps}
        />
      </ul>
    </>
  );
}

export { FileTree };
