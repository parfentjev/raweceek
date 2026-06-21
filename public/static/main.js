const messageContainer = document.getElementById("message");
const countdownTemplate = document.getElementById("countdown-item-template").content.firstElementChild;
const countdownsContainer = document.getElementById("countdowns");
const nextSessionJsonContainer = document.getElementById("next-session-json");

const error = (message) => {
    messageContainer.className = "error";
    messageContainer.textContent = message;
};

/** I don't need this, but I want to keep it just in case.
const rainbow = (element) => {
  const input = element.textContent;
  element.textContent = "";
  
  let color = Math.floor(Math.random() * 7);
  for (const char of input) {
    element.innerHTML += `<span class="color-${color++ % 6}">${char}</span>`;
  }
};
*/

const run = () => {
    fetch("/api/status").then((response) => {
      if (!response.ok) {
        throw Error(`Oi, mate. I tried to call the bloody API, but that fookin' wanker returned: ${response.status}`);
      }

      return response.json();
    }).then((data) => {
      if (!data || data.raceWeek === undefined || !data.nextSession || !data.nextSession.countdowns) {
        throw Error('Oi, mate. The bloody service returned some garbage. I can't work with that, mate.');
      }

      const raceWeekBlockId = data.raceWeek ? "race-week" : "not-race-week";
      document.getElementById(raceWeekBlockId).hidden = false;

      for (const countdown of data.nextSession.countdowns) {
        const item = countdownTemplate.cloneNode(true);
        const colorId = Math.floor(Math.random() * 7);
        item.className = `color-${colorId}`;
        item.textContent = countdown.value;

        countdownsContainer.append(item);
      }

      nextSessionJsonContainer.textContent = JSON.stringify(data);

      messageContainer.hidden = true;
    }).catch((message) => error(message));
};

run();
