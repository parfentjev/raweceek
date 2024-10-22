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
                        href="https://www.formula1.com/en/racing/2024"
                        target="_blank"
                        rel="noreferrer"
                    >
                        📅calendar
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
            </ul>
        </nav>
    )
}

export default Navigation
