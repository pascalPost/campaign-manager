'use client';

import { Button } from "@/components/ui/button";
import { OpenFolderDialog } from "@/wailsjs/go/main/App";

export default function Home() {

    function onClick() {
        console.log("Open Project");

        OpenFolderDialog().then((dir :string) => {
            console.log(`Selected Folder: ${dir}`)
        });
    }

    return (
        <>
            <h1>Home (Dashboard)</h1>
        <div className="flex justify-center items-center h-screen">
            <Button onClick={onClick}>Open Project</Button>
        </div>
        </>
    )
}
