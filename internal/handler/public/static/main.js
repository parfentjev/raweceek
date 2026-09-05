function throwErr(details) {
  const msg = "Oi, mate. Something went wrong: {}";

  throw Error(msg.replace("{}", details));
}

const messageContainer = document.getElementById("message");
const upcomingSessionsContainer = document.getElementById("upcoming-sessions");
const nextSessionJsonContainer = document.getElementById("next-session-json");

const sessionTemplate = document.getElementById("session-details-template").content.firstElementChild;
const countdownTemplate = document.getElementById("countdown-item-template").content.firstElementChild;

const error = (message) => {
  messageContainer.className = "error";
  messageContainer.textContent = message;
};

const run = () => {
  fetch("/api/v2/status")
    .then((response) => {
      if (!response.ok) {
        throwErr(`unexpected status code ${response.status}`);
      }

      return response.json();
    })
    .then((data) => {
      if (data === undefined) {
        throwErr("server returned no data");
      }

      const { upcoming_sessions: upcoming, race_week: raceWeek } = data;
      if (upcoming === undefined || raceWeek === undefined) {
        throwErr("server returned some garbage");
      }

      const raceWeekBlockId = raceWeek ? "race-week" : "not-race-week";
      document.getElementById(raceWeekBlockId).hidden = false;

      for (const session of upcoming) {
        const sessionContainer = sessionTemplate.cloneNode(true);
        sessionContainer.querySelector("h2").textContent = session.summary;

        for (const countdown of session.countdowns) {
          const countdownContainer = countdownTemplate.cloneNode(true);
          const colorId = Math.floor(Math.random() * 7);
          countdownContainer.className = `color-${colorId}`;
          countdownContainer.textContent = countdown.value;

          const countdownsContainer = sessionContainer.querySelector("ul");
          countdownsContainer.append(countdownContainer);
        }

        upcomingSessionsContainer.append(sessionContainer);
      }

      nextSessionJsonContainer.textContent = JSON.stringify(data);
      messageContainer.hidden = true;
    })
    .catch((message) => error(message));
};

run();
