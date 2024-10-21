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
                        href="https://codeberg.org/parfentjev/raweceek/src/branch/master/backend/spec/swagger.yaml"
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
        </nav>
    )
}

export default Navigation
