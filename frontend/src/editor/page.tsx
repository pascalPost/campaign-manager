import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "@/components/ui/resizable";
import { FileTree, FileTreeRoot } from "@/components/file-tree.tsx";
import { useState } from "react";

const data: FileTreeRoot = {
  id: "/root",
  name: "root",
  content: [
    { id: "/root/file_one.txt", type: "file", name: "file_one.txt" },
    {
      id: "/root/folder one",
      type: "folder",
      name: "folder one",
      content: [
        { id: "/root/folder one/file_one", type: "file", name: "file_one" },
        { id: "/root/folder one/file_two", type: "file", name: "file_two" },
      ],
    },
    {
      id: "/root/folder_two",
      type: "folder",
      name: "folder_two",
      content: [
        { id: "/root/folder_two/file one", type: "file", name: "file one" },
        { id: "/root/folder_two/file_two", type: "file", name: "file_two" },
      ],
    },
  ],
};

export function EditorPage() {
  const [selectedFile, setSelectedFile] = useState<string>("");

  return (
    <div className="flex h-full flex-col gap-4">
      <h1 className="scroll-m-20 text-2xl font-semibold tracking-tight">
        Editor
      </h1>
      <ResizablePanelGroup
        direction="horizontal"
        className="w-full rounded-md border"
      >
        <ResizablePanel defaultSize={25}>
          <FileTree
            data={data}
            selectedFileProps={{
              selectedFile: selectedFile,
              handleChangeSelectedFile: (fileId: string) => {
                setSelectedFile(fileId);
              },
            }}
          />
        </ResizablePanel>
        <ResizableHandle withHandle />
        <ResizablePanel>{selectedFile}</ResizablePanel>
      </ResizablePanelGroup>
    </div>
  );
}
