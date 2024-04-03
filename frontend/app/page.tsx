'use client';

import {Button} from "@/components/ui/button";
import {OpenFolderDialog} from "@/wailsjs/go/main/App";
import {Home} from "lucide-react";
import {Card} from "@/components/ui/card";
import {ProjectTable} from "@/app/project-table";

export default function HomePage() {

    // function onClick() {
    //     console.log("Open Project");
    //
    //     OpenFolderDialog().then((dir: string) => {
    //         console.log(`Selected Folder: ${dir}`)
    //     });
    // }

    return (
        <>
            <h1 className="scroll-m-20 text-2xl font-semibold tracking-tight">Dashboard</h1>
            <div className="flex gap-5 mt-4">
                <Button className="w-full"><a href="/project/create">Create New Project</a></Button>
            </div>
            <div className="mt-4">
                <ProjectTable/>
            </div>
        </>
    )
}
