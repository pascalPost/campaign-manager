import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "@/components/ui/resizable";
import { FileTree, FileTreeRoot } from "@/components/file-tree.tsx";
import { useEffect, useState } from "react";
import { Textarea } from "@/components/ui/textarea.tsx";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { client } from "@/lib/api/client.ts";
import { Button } from "@/components/ui/button.tsx";

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

async function getFile(signal?: AbortSignal): Promise<string> {
  const { data } = await client.GET("/file/{filePath}", {
    params: {
      path: {
        filePath: "test.txt",
      },
    },
    parseAs: "text",
    signal,
  });
  return data || "";
}

async function putFile(text: string): Promise<void> {
  await client.PUT("/file/{filePath}", {
    params: {
      path: {
        filePath: "test.txt",
      },
    },
    body: text,
    // work-around: deactivate default json.stringify, see https://github.com/OpenAPITools/openapi-generator/issues/7083
    bodySerializer(body) {
      return body;
    },
    parseAs: "text",
  });
}

export function EditorPage() {
  const [selectedFile, setSelectedFile] = useState<string>("");
  const [editorText, setEditorText] = useState<string>("");

  const query = useQuery({
    queryKey: ["getFile"],
    queryFn: ({ signal }) => getFile(signal),
    staleTime: Infinity,
    cacheTime: Infinity,
  });

  useEffect(() => {
    if (query.isSuccess) {
      setEditorText(query.data || "");
    }
  }, [query.isSuccess, query.data]);

  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (fileContent: string) => putFile(fileContent),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["getFile"] }),
  });

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
        <ResizablePanel>
          {/*{selectedFile}*/}
          <div className="flex h-full flex-col gap-2">
            <Textarea
              value={editorText}
              onChange={(e) => setEditorText(e.target.value)}
              className="h-full resize-none border-none focus-visible:ring-0"
              id="editor-textarea"
            />
            <div className="flex flex-row gap-2">
              <Button
                variant="secondary"
                className="w-full"
                onClick={() => {
                  setEditorText(query.data || "");
                }}
              >
                Cancel
              </Button>
              <Button
                className="w-full"
                onClick={() => mutation.mutate(editorText)}
              >
                Save
              </Button>
            </div>
          </div>
        </ResizablePanel>
      </ResizablePanelGroup>
    </div>
  );
}
