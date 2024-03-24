import {Button} from "@/components/ui/button";
import {Input} from "@/components/ui/input";

export default function CreateProjectPage() {
    return (
        <>
            <h1>Create Project</h1>
            <div>
                Name: <Input type="text" />
            </div>
            <Button>Open CSV</Button>
            <input type="file" />
            <Button>Start</Button>
            <Button>Cancel</Button>
        </>
    )
}