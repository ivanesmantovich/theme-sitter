interface CodeProps {
    text: string
}

export default function Code({ text }: CodeProps) {
    return (
        <div className="overflow-y-hidden whitespace-pre-wrap">{text}</div>
    )
}
