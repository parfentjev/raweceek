import { FC, useEffect, useState } from 'react'
import { sessionApi } from '../api/api'
import ChristmasGarland from '../components/ChristmasGarland'
import WelcomeBox from '../components/WelcomeBox'
import Navigation from '../components/Navigation'
import Counter from '../components/Counter'
import { SessionDto } from '../api/codegen'
import Banners from '../components/Banners'

const WelcomePage: FC = () => {
    const [session, setSession] = useState<SessionDto>()

    useEffect(() => {
        sessionApi()
            .sessionsNextGet()
            .then((session) => setSession(session))
            .catch((e) => console.log('err response', e))
    }, [])

    return (
        <>
            <ChristmasGarland />
            <Banners />
            <Navigation />
            <WelcomeBox session={session} />
            <Navigation />
            <Counter />
        </>
    )
}

export default WelcomePage
