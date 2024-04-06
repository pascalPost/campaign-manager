import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { catchError } from "@/lib/utils.ts";

const settingsSchema = z.object({
  workingDir: z.string().min(1),
});

export default function SettingsPage() {
  const form = useForm<z.infer<typeof settingsSchema>>({
    resolver: zodResolver(settingsSchema),
    defaultValues: {
      workingDir: "", // load from DB
    },
  });

  function onSubmit(values: z.infer<typeof settingsSchema>) {
    console.log(values);
  }

  function onReset() {
    form.reset();
  }

  return (
    <div className="px-0 md:px-8">
      <h1 className="scroll-m-20 text-2xl font-semibold tracking-tight">
        Settings
      </h1>
      <div className="w-full py-4 md:max-w-2xl">
        <Form {...form}>
          <form
            onSubmit={(event) => catchError(form.handleSubmit(onSubmit)(event))}
            onReset={onReset}
            className="space-y-8"
          >
            <FormField
              control={form.control}
              name="workingDir"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Working Directory</FormLabel>
                  <FormControl>
                    <Input placeholder="/work/data" type="text" {...field} />
                  </FormControl>
                  <FormDescription>
                    The working directory (on the server) where the project
                    files are stored.
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
            <div className="flex gap-4">
              <Button type="submit">Save</Button>
              <Button type="reset" variant="secondary">
                Cancel
              </Button>
            </div>
          </form>
        </Form>
      </div>
    </div>
  );
}
