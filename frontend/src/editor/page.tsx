import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "@/components/ui/resizable";
import { FileTree, FileTreeRoot } from "@/components/file-tree.tsx";
import { useEffect, useState } from "react";
import { Editor } from "@/components/editor.tsx";
import { useQuery } from "react-query";
import { client } from "@/lib/api/client.ts"; // const data: FileTreeRoot = {

async function getFileTree(signal?: AbortSignal) {
  const fileTree: FileTreeRoot = {
    id: "/",
    name: "/",
    content: [],
  };

  const { data, error } = await client.GET("/fileTree", {
    signal,
  });

  if (error) throw error;
  if (data) {
    data.map((entry) => {
      if (entry.isDir) {
        fileTree.content.push({
          id: `/${entry.name}`,
          type: "folder",
          name: entry.name,
          content: [],
        });
      } else {
        fileTree.content.push({
          id: `/${entry.name}`,
          type: "file",
          name: entry.name,
        });
      }
    });
  }

  return fileTree;
}

export function EditorPage() {
  const [fileTree, setFileTree] = useState<FileTreeRoot>({
    id: "/",
    name: "/",
    content: [],
  });
  const [selectedFile, setSelectedFile] = useState<string>("");

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
