import { FC, useEffect, useState } from 'react'
import { sessionApi } from '../api/api'
import { toast } from 'react-toastify'
import { getWeek } from 'date-fns'
import Notifications from '../components/Notifications'
import ChristmasGarland from '../components/ChristmasGarland'
import WelcomeBox from '../components/WelcomeBox'
import Navigation from '../components/Navigation'

const WelcomePage: FC = () => {
    const [raceWeek, setRaceWeek] = useState<boolean>()

    useEffect(() => {
        sessionApi()
            .sessionsNextGet()
            .then((session) => {
                const sessionTime = new Date(session.startTime)
                const currentTime = new Date()

                const sameWeek =
                    getWeek(sessionTime, { weekStartsOn: 1 }) ===
                    getWeek(currentTime, { weekStartsOn: 1 })

                if (sameWeek) {
                    setRaceWeek(true)
                } else {
                    setRaceWeek(false)
                }
            })
            .catch(() => toast.error('Something went wrong!!'))
    }, [])

    return (
        <>
            <Notifications />
            <ChristmasGarland />
            <Navigation />
            <WelcomeBox raceWeek={raceWeek} />
            <Navigation />
        </>
    )
}

export default WelcomePage
