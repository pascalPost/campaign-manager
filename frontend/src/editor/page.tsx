import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "@/components/ui/resizable";
import {
  File,
  FileTree,
  FileTreeRoot,
  Folder,
} from "@/components/file-tree.tsx";
import { useEffect, useState } from "react";
import { Editor } from "@/components/editor.tsx";
import { useQuery } from "react-query";
import { client } from "@/lib/api/client.ts"; // const data: FileTreeRoot = {

async function getFileTree(signal?: AbortSignal) {
  const fileTree: FileTreeRoot = {
    path: "/",
    content: new Array<File | Folder>(),
  };

  const { data, error } = await client.GET("/fileTree", {
    signal,
  });

  if (error) throw error;

  if (data) {
    data.map((entry) => {
      if (entry.isDir) {
        fileTree.content!.push({
          type: "folder",
          path: `/${entry.name}`,
        });
      } else {
        fileTree.content!.push({
          type: "file",
          path: `/${entry.name}`,
        });
      }
    });
  }

  return fileTree;
}

export function EditorPage() {
  const [fileTree, setFileTree] = useState<FileTreeRoot>({
    path: "/",
  });
  const [selectedFile, setSelectedFile] = useState<string | undefined>(
    undefined,
  );

  const query = useQuery({
    queryKey: ["getFileTree"],
    queryFn: ({ signal }) => getFileTree(signal),
    staleTime: Infinity,
    cacheTime: Infinity,
  });

  useEffect(() => {
    if (query.isSuccess && query.data) {
      setFileTree(query.data);
    }
  }, [query.isSuccess, query.data]);

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
            data={fileTree}
            selectedFileProps={{
              selectedFile: selectedFile,
              handleChangeSelectedFile: (fileId: string) => {
                setSelectedFile(fileId);
              },
            }}
          />
        </ResizablePanel>
        <ResizableHandle withHandle />
        <ResizablePanel>
          <Editor filePath={selectedFile} />
        </ResizablePanel>
      </ResizablePanelGroup>
    </div>
  );
}
