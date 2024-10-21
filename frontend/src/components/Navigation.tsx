import { FC } from 'react'

const Navigation: FC = () => {
    return (
        <nav>
            <ul>
                <li>
                    <a
                        href="http://users4.smartgb.com/g/g.php?a=s&i=g44-82594-15"
                        target="_blank"
                        rel="noreferrer"
                    >
                        📚guest book
                    </a>
                </li>
                <li>
                    <a
                        href="https://codeberg.org/parfentjev/raweceek"
                        target="_blank"
                        rel="noreferrer"
                    >
                        📜api
                    </a>
                </li>
                <li>
                    <a
                        href="https://codeberg.org/parfentjev/raweceek"
                        target="_blank"
                        rel="noreferrer"
                    >
                        👩‍💻source
                    </a>
                </li>
            </ul>
            {process.env.NODE_ENV === 'production' && (
                <div>
                    <a href="https://www.free-website-hit-counter.com">
                        <img
                            src="https://www.free-website-hit-counter.com/zc.php?d=9&id=532&s=1"
                            alt="Free Website Hit Counter"
                        />
                    </a>
                </div>
            )}
        </nav>
    )
}

export default Navigation
