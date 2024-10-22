import { FC, useEffect, useState } from 'react'
import { sessionApi } from '../api/api'
import { toast } from 'react-toastify'
import Notifications from '../components/Notifications'
import ChristmasGarland from '../components/ChristmasGarland'
import WelcomeBox from '../components/WelcomeBox'
import Navigation from '../components/Navigation'
import Counter from '../components/Counter'
import { SessionDto } from '../api/codegen'

const WelcomePage: FC = () => {
    const [session, setSession] = useState<SessionDto>()

    useEffect(() => {
        sessionApi()
            .sessionsNextGet()
            .then((session) => setSession(session))
            .catch(() => toast.error('Something went wrong!!'))
    }, [])

    return (
        <>
            <Notifications />
            <ChristmasGarland />
            <Navigation />
            <WelcomeBox session={session} />
            <Navigation />
            <Counter />
        </>
    )
}

export default WelcomePage
