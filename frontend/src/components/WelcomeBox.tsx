import { FC } from 'react'
import RainbowText from './RainbowText'
import { SessionDto } from '../api/codegen'
import { getWeek } from 'date-fns'

interface Props {
    session?: SessionDto
}

const RACE_WEEK_TRUE = 'I know that this week is a race week!'
const RACE_WEEK_FALSE = 'This is not a race week! :('
const RACE_WEEK_UNDEFINED = "I dont't know if this week is a race week..."

const WelcomeBox: FC<Props> = ({ session }) => {
    let message: any = RACE_WEEK_UNDEFINED

    if (session) {
        const sessionTime = new Date(session.startTime)
        const currentTime = new Date()

        const sameWeek =
            getWeek(sessionTime, { weekStartsOn: 1 }) ===
            getWeek(currentTime, { weekStartsOn: 1 })

        if (sameWeek) {
            message = (
                <h1>
                    <RainbowText>{RACE_WEEK_TRUE}</RainbowText>
                </h1>
            )
        } else {
            message = RACE_WEEK_FALSE
        }
    }

    return (
        <div id="welcome-box">
            <p>
                <img src="media/welcome.png" alt="Welcome!" />
            </p>
            <p>{message}</p>
            <p>
                <img src="media/kramer-mind-blown.gif" alt="Mind blowing!" />
            </p>
            <p>'But how??' you ask.</p>
            <p>
                <img src="media/super-easy.png" alt="Super easy" />
            </p>
            <p>
                I say. Thanks to this <RainbowText>majestic</RainbowText>...
            </p>
            <p>
                <img src="media/rest-api-service.png" alt="REST API Service" />
            </p>
            <p>
                I only need to make <RainbowText>one request</RainbowText> to
                know when is the next racing session!
            </p>
            <p>
                <img
                    className="dancing-image"
                    src="media/is-that-curl.png"
                    alt="Is that CURL?"
                />
            </p>
            <p>
                <RainbowText>YES IT IS!!!</RainbowText>
            </p>
            <code>$ curl https://raweceek.eu/api/sessions/next</code>
            {session && <code>{JSON.stringify(session)}</code>}
            <p>
                <img src="media/wow.png" alt="Wow!" />
            </p>
            <p>
                <img src="media/horses.gif" alt="Horses" />
            </p>
        </div>
    )
}

export default WelcomeBox
