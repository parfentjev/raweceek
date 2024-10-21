import { FC } from 'react'
import RainbowText from './RainbowText'

interface Props {
    raceWeek?: boolean
}

const RACE_WEEK_TRUE = 'I know that this week is a race week!'
const RACE_WEEK_FALSE = 'This is not a race week! :('
const RACE_WEEK_UNDEFINED = "I dont't know if this week is a race week..."

const WelcomeBox: FC<Props> = ({ raceWeek }) => {
    let message

    switch (raceWeek) {
        case true:
            message = (
                <h1>
                    <RainbowText>{RACE_WEEK_TRUE}</RainbowText>
                </h1>
            )
            break
        case false:
            message = RACE_WEEK_FALSE
            break
        default:
            message = RACE_WEEK_UNDEFINED
            break
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
                I say. Thanks to this <RainbowText>fantastic</RainbowText>...
            </p>
            <p>
                <img src="media/rest-api-service.png" alt="REST API Service" />
            </p>
            <p>
                I only need to make <RainbowText>one request</RainbowText> to
                know when is the next race!
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
            <code>
                {`{
    summary:
        '\uD83C\uDFC1 FORMULA 1 PIRELLI UNITED STATES GRAND PRIX 2024 - Race',
    location: 'United States',
    startTime: '2024-10-20T19:00Z',
}`}
            </code>
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
