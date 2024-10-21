import { Configuration, SessionApi } from "./codegen"

const configuration = () => {
    return new Configuration({
        basePath: process.env.REACT_APP_SERVICE_URL,
    })
}

export const sessionApi = () => {
    return new SessionApi(configuration())
}
