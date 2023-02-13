import './Code.css'

interface CodeProps {
    text: string
}

export default function Code({ text }: CodeProps) {
    return (
        <div className="code">{text}</div>
    )
}
