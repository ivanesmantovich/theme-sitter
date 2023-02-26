interface CodeProps {
    text: string
}

export default function Code({ text }: CodeProps) {
    return (
        <div className="whitespace-pre">{text}</div>
    )
}
