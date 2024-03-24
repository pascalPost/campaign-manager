'use client';

import {Button} from "@/components/ui/button";
import {OpenFolderDialog} from "@/wailsjs/go/main/App";
import {Home} from "lucide-react";
import {Card} from "@/components/ui/card";

export default function HomePage() {

    function onClick() {
        console.log("Open Project");

        OpenFolderDialog().then((dir: string) => {
            console.log(`Selected Folder: ${dir}`)
        });
    }

    return (
        <>
            <h1>Home (Dashboard)</h1>
            <div>
                <Button><a href="/createProject">Create Project</a></Button>
                <Button onClick={onClick}>Open Project</Button>
            </div>
            <Card>
                Project
                <Button><a href="/project">Open</a></Button>
            </Card>
        </>
    )
}
