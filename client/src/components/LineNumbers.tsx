import { useEffect, useState } from "react";
import './LineNumbers.css'

interface LineNumbersProps {
    numberOfLines: number
}

export default function LineNumbers({ numberOfLines }: LineNumbersProps) {
    const [lineNumbers, setLineNumbers] = useState<JSX.Element[]>([])

    function computeLineNumbers(): void {
        const lineNumbers: JSX.Element[] = [];
        for (let i = 0; i < numberOfLines; i++) {
            lineNumbers.push(<span key={i}></span>);
        }
        setLineNumbers(lineNumbers);
    }

    useEffect(() => {
        computeLineNumbers();
    }, [numberOfLines]);

    return (
        <div className="line-numbers">
            {lineNumbers}
        </div>
    )
}
