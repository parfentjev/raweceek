"use strict";
const find = {
    elementBy(args) {
        const e = document.getElementById(args.id);
        if (e == null)
            throw Error(`element not found by id: ${args.id}`);
        return e;
    },
    templateBy(args) {
        const e = document.getElementById(args.id);
        if (e == null)
            throw Error(`template not found by id: ${args.id}`);
        return e;
    },
    childBy(args) {
        const e = args.parent.querySelector(args.selector);
        if (e == null)
            throw Error(`child not found by selector: ${args.selector}`);
        return e;
    },
};
const clone = (template) => {
    const child = template.content.firstElementChild;
    if (child == null)
        throw Error(`invalid HTML template`);
    return child.cloneNode(true);
};
const errorMessage = "Oi, mate, this isn't working: {}";
const messageContainer = find.elementBy({ id: "message" });
const run = async () => {
    const upcomingSessionsContainer = find.elementBy({ id: "upcoming-sessions" });
    const nextSessionJsonContainer = find.elementBy({ id: "next-session-json" });
    const sessionTemplate = find.templateBy({ id: "session-details-template" });
    const countdownTemplate = find.templateBy({ id: "countdown-item-template" });
    const parseResponse = (response) => {
        if (!response.ok) {
            throw Error(`unexpected status code ${response.status}`);
        }
        return response.json();
    };
    await fetch("/api/v2/status")
        .then((response) => parseResponse(response))
        .then((data) => {
        const { upcoming_sessions, race_week } = data;
        const raceWeekBlockId = race_week ? "race-week" : "not-race-week";
        find.elementBy({ id: raceWeekBlockId }).hidden = false;
        for (const session of upcoming_sessions) {
            const sessionContainer = clone(sessionTemplate);
            const summaryHeader = find.childBy({ parent: sessionContainer, selector: "h2" });
            const startTimeHeader = find.childBy({ parent: sessionContainer, selector: "h3" });
            if (summaryHeader == null || startTimeHeader == null)
                throw Error(`Missing session data`);
            summaryHeader.textContent = session.summary;
            startTimeHeader.textContent = session.start_time;
            for (const countdown of session.countdowns) {
                const countdownContainer = clone(countdownTemplate);
                countdownContainer.textContent = countdown.value;
                const countdownsContainer = find.childBy({ parent: sessionContainer, selector: "ul" });
                if (countdownContainer == null)
                    throw Error();
                countdownsContainer.append(countdownContainer);
            }
            upcomingSessionsContainer.append(sessionContainer);
        }
        nextSessionJsonContainer.textContent = JSON.stringify(data);
        messageContainer.hidden = true;
    });
};
run().catch((error) => {
    messageContainer.className = "error";
    messageContainer.textContent = errorMessage.replace("{}", error);
});
