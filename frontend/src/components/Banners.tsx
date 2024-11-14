import { FC, useEffect, useState } from 'react'

const Banners: FC = () => {
    const [filename, setFilename] = useState('')

    useEffect(() => {
        fetch('./banners.txt')
            .then((data) => data.text())
            .then((banners) => {
                const lines = banners.split('\n')
                const index = Math.floor(Math.random() * lines.length)
                setFilename(lines[index])
            })
            .catch((e) => console.log(e))
    }, [])

    return (
        <div id="banner">
            {filename && (
                <a
                    href="https://www.youtube.com/watch?v=zWaymcVmJ-A"
                    target="_blank"
                    rel="noreferrer"
                >
                    <img src={`./banners/${filename}`} alt="Banner" />
                </a>
            )}
        </div>
    )
}

export default Banners
