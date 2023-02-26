import { useEffect, useState } from "react";
import './CodeWindow.css'
import Code from './Code';
import LineNumbers from './LineNumbers';

export default function CodeWindow() {
    const [numberOfLines, setNumberOfLines] = useState(0);
    const text = "code line 1 Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.\ncode line 2\ncode line 3\ncode line 4\ncode line 5\ncode line 6\ncode line 7\ncode line 8\ncode line 9\ncode line 10\ncode line 11\ncode line 12\ncode line 13\ncode line 14\ncode line 15\ncode line 16\ncode line 17\ncode line 18\ncode line 19\ncode line 20";

    function countTheLines(): void {
        setNumberOfLines(text.split('\n').length);
    }

    useEffect(() => {
        countTheLines();
    }, []);

    return (
        <div className='code-window w-3/4 h-96 overflow-auto inline-flex m-5 p-2 text-lg gap-3 border border-solid'>
            <LineNumbers numberOfLines={numberOfLines}></LineNumbers>
            <Code text={text}></Code>
        </div>
    )
}
