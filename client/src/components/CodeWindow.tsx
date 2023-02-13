import { useEffect, useState } from "react";
import './CodeWindow.css'
import Code from './Code';
import LineNumbers from './LineNumbers';

export default function CodeWindow() {
    const [numberOfLines, setNumberOfLines] = useState(0);
    const text = "code line 1\ncode line 2\ncode line 3";

    function countTheLines(): void {
        setNumberOfLines(text.split('\n').length);
    }

    useEffect(() => {
        countTheLines();
    }, []);

    return (
        <div className='code-window'>
            <LineNumbers numberOfLines={numberOfLines}></LineNumbers>
            <Code text={text}></Code>
        </div>
    )
}
