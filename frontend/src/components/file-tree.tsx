import { ReactElement, useState } from "react";
import { ChevronDown, ChevronRight } from "lucide-react";

type File = {
  id: string;
  type: "file";
  name: string;
};

type Folder = {
  id: string;
  type: "folder";
  name: string;
  content: Array<File | Folder>;
};

type FileTreeRoot = {
  id: string;
  name: string;
  content: Array<File | Folder>;
};

type SelectedFileProps = {
  selectedFile: string;
  handleChangeSelectedFile: (fileId: string) => void;
};

function FileTreeEntries(
  content: Array<File | Folder>,
  selectedFileProps: SelectedFileProps,
): ReactElement {
  return (
    <ul className="px-4">
      {content.map((e: File | Folder) => {
        switch (e.type) {
          case "file":
            return (
              <FileTreeFile
                key={e.id}
                file={e}
                selectedFileProps={selectedFileProps}
              />
            );
          case "folder":
            return (
              <FileTreeFolder
                key={e.id}
                name={e.name}
                content={e.content}
                selectedFileProps={selectedFileProps}
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
  return (
    <li className="flex flex-row">
      <div className="w-5" />
      <div
        onClick={() => selectedFileProps.handleChangeSelectedFile(file.id)}
        className="hover:cursor-pointer"
      >
        {file.name}
      </div>
    </li>
  );
}

function FileTreeFolder({
  name,
  content,
  selectedFileProps,
}: {
  name: string;
  content: Array<File | Folder>;
  selectedFileProps: SelectedFileProps;
}): ReactElement {
  const [isExpanded, setExpanded] = useState<boolean>(false);

  if (!isExpanded) {
    return (
      <>
        <li className="flex flex-row items-center">
          <ChevronRight
            className="mt-1 h-4 w-5 hover:cursor-pointer"
            onClick={() => setExpanded(true)}
          />
          <div>{name}</div>
        </li>
      </>
    );
  }

  return (
    <>
      <li className="flex flex-row items-center">
        <ChevronDown
          className="mt-1 h-4 w-5 hover:cursor-pointer"
          onClick={() => setExpanded(false)}
        />
        {name}
      </li>
      {FileTreeEntries(content, selectedFileProps)}
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
  return (
    <>
      <ul className="px-1">
        <FileTreeFolder
          key={data.id}
          name={data.name}
          content={data.content}
          selectedFileProps={selectedFileProps}
        />
      </ul>
    </>
  );
}

export { FileTree, type FileTreeRoot, type Folder, type File };
