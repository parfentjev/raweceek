import { FC } from 'react'

const Counter: FC = () => {
    return (
        <div id="counter">
            {process.env.NODE_ENV === 'production' && (
                <a href="https://www.free-website-hit-counter.com">
                    <img
                        src="https://www.free-website-hit-counter.com/zc.php?d=9&id=532&s=1"
                        alt="Free Website Hit Counter"
                    />
                </a>
            )}
        </div>
    )
}

export default Counter
