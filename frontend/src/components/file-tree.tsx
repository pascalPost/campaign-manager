import { ReactElement, useState } from "react";
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
  content?: Array<File | Folder>;
};

type FileTreeRoot = {
  path: string;
  content?: Array<File | Folder>;
};

type SelectedFileProps = {
  selectedFile?: string;
  handleChangeSelectedFile: (fileId: string) => void;
};

function FileTreeEntries(
  isExpandedMap: Map<string, boolean>,
  setExpandedMap: (path: string, newState: boolean) => void,
  selectedFileProps: SelectedFileProps,
  content?: Array<File | Folder>,
): ReactElement {
  return (
    <ul className="px-4">
      {(content || []).map((e: File | Folder) => {
        switch (e.type) {
          case "file":
            return (
              <FileTreeFile
                key={e.path}
                file={e}
                selectedFileProps={selectedFileProps}
              />
            );
          case "folder":
            return (
              <FileTreeFolder
                key={e.path}
                path={e.path}
                content={e.content}
                selectedFileProps={selectedFileProps}
                isExpandedMap={isExpandedMap}
                setExpandedMap={setExpandedMap}
              />
            );
        }
      })}
    </ul>
  );
}

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

// async function getFileTree(path: string, signal?: AbortSignal) {
//   const {data, error} = await client.GET(`/fileTree/{filePath}`, {
//     params: {
//       path: {
//         filePath: path,
//       },
//     },
//     signal,
//   });
//
//   if(error) throw error;
//
//   if(data) {
//
//   }
// }

function FileTreeFolder({
  path,
  content,
  isExpandedMap,
  setExpandedMap,
  selectedFileProps,
}: {
  path: string;
  content?: Array<File | Folder>;
  isExpandedMap: Map<string, boolean>;
  setExpandedMap: (path: string, state: boolean) => void;
  selectedFileProps: SelectedFileProps;
}): ReactElement {
  const folderName = path.split("/").pop() || path;
  const isExpanded = isExpandedMap.get(path) || false;

  if (!isExpanded) {
    return (
      <>
        <li className="flex flex-row items-center">
          <ChevronRight
            className="mt-1 h-4 w-5 hover:cursor-pointer"
            onClick={() => setExpandedMap(path, true)}
          />
          <div>{folderName}</div>
        </li>
      </>
    );
  }

  // if (content === undefined) {
  //   const query = useQuery({
  //     queryKey: [`getFileTree_${path}`],
  //     queryFn: getFileTree(path),
  //   });
  //   console.log(`request file tree for root ${path}`);
  // }

  return (
    <>
      <li className="flex flex-row items-center">
        <ChevronDown
          className="mt-1 h-4 w-5 hover:cursor-pointer"
          onClick={() => setExpandedMap(path, false)}
        />
        {folderName}
      </li>
      {FileTreeEntries(
        isExpandedMap,
        setExpandedMap,
        selectedFileProps,
        content,
      )}
    </>
  );
}

function FileTree({
  data,
  selectedFileProps,
}: {
  data: FileTreeRoot;
  selectedFileProps: SelectedFileProps;
}) {
  const [isExpandedMap, setExpandedMap] = useState<Map<string, boolean>>(
    new Map(),
  );

  return (
    <>
      <ul className="px-1">
        <FileTreeFolder
          key={data.path}
          path={data.path}
          content={data.content}
          isExpandedMap={isExpandedMap}
          setExpandedMap={(path: string, newState: boolean) => {
            setExpandedMap(
              new Map<string, boolean>(isExpandedMap.set(path, newState)),
            );
          }}
          selectedFileProps={selectedFileProps}
        />
      </ul>
    </>
  );
}

export { FileTree, type FileTreeRoot, type Folder, type File };
