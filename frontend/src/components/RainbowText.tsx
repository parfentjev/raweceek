import { FC, ReactNode } from 'react'

interface Props {
    children: ReactNode
}

const RainbowText: FC<Props> = ({ children }) => {
    const colors = [
        '#FF0000',
        '#FF7F00',
        '#FFFF00',
        '#00FF00',
        '#0000FF',
        '#4B0082',
        '#EE82EE',
    ]

    return (
        <>
            {children
                ?.toString()
                .split('')
                .map((letter, index) => (
                    <span
                        key={index}
                        style={{ color: colors[index % colors.length] }}
                    >
                        {letter}
                    </span>
                ))}
        </>
    )
}

export default RainbowText
