import { Textarea } from "@/components/ui/textarea.tsx";
import { Button } from "@/components/ui/button.tsx";
import { useEffect, useState } from "react";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { client } from "@/lib/api/client.ts";

async function getFile(
  filePath: string,
  signal?: AbortSignal,
): Promise<string> {
  const { data } = await client.GET("/file/{filePath}", {
    params: {
      path: {
        filePath: filePath.replace(/^\/+/g, ""), // remove leading /
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

function Editor({ filePath }: { filePath: string }) {
  const [editorText, setEditorText] = useState<string>("");

  const query = useQuery({
    queryKey: [`getFile${filePath}`],
    queryFn: ({ signal }) => getFile(filePath, signal),
  });

  useEffect(() => {
    if (query.isSuccess) {
      setEditorText(query.data || "");
    }
  }, [query.isSuccess, query.data]);

  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (fileContent: string) => putFile(fileContent),
    onSuccess: () =>
      queryClient.invalidateQueries({ queryKey: [`getFile${filePath}`] }),
  });

  return (
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
        <Button className="w-full" onClick={() => mutation.mutate(editorText)}>
          Save
        </Button>
      </div>
    </div>
  );
}

export { Editor };
