'use client';

import {Inter} from 'next/font/google'
import './globals.css'
import {cn} from "@/lib/utils";

const inter = Inter({subsets: ['latin'], variable: "--font-sans"})

export default function RootLayout({
                                       children,
                                   }: {
    children: React.ReactNode
}) {
    return (
        <html lang="en" suppressHydrationWarning={true}>
        <body className={cn(
            "min-h-screen bg-background font-sans antialiased ml-2 mr-2",
            inter.variable
        )}>
        {children}
        </body>
        </html>
    )
}
