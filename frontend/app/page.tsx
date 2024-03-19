'use client';

import {Button} from "@/components/ui/button";
import {OpenFolderDialog} from "@/wailsjs/go/main/App";
import {NavBreadcrumb} from "@/components/nav-breadcrumb";
import {Home} from "lucide-react";

export default function HomePage() {

    function onClick() {
        console.log("Open Project");

        OpenFolderDialog().then((dir: string) => {
            console.log(`Selected Folder: ${dir}`)
        });
    }

    return (
        <>
            <NavBreadcrumb entries={[
                {
                    href: "/",
                    label: (
                        <div className="inline-flex items-center gap-1">
                            <Home size="1em"/>
                            Home
                        </div>
                    ),
                },
            ]}/>
            <h1>Home (Dashboard)</h1>
            <div className="flex justify-center items-center h-screen">
                <Button><a href="/createProject">Create Project</a></Button>
                <Button onClick={onClick}>Open Project</Button>
            </div>
        </>
    )
}
