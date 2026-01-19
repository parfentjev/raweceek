/*
    I couldn't be arsed to write this myself, so I asked claude.

    This converts an ics file that is usually available on f1.com to insert statements.

    Usage: bun ics.js input.ics output.sql
*/

import { readFileSync, writeFileSync } from "fs";

function parseICS(fileContent) {
  const events = [];
  const lines = fileContent.split(/\r?\n/);

  let currentEvent = null;
  let currentField = null;

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i];

    if (line === "BEGIN:VEVENT") {
      currentEvent = {};
    } else if (line === "END:VEVENT") {
      if (currentEvent) {
        events.push(currentEvent);
        currentEvent = null;
      }
    } else if (currentEvent) {
      // Handle line continuation (lines starting with space or tab)
      if (line.match(/^[ \t]/)) {
        if (currentField) {
          currentEvent[currentField] += line.substring(1);
        }
        continue;
      }

      const colonIndex = line.indexOf(":");
      if (colonIndex === -1) continue;

      const fieldName = line.substring(0, colonIndex);
      const fieldValue = line.substring(colonIndex + 1);

      // Extract the base field name (before any parameters like ;TZID=)
      const baseFieldName = fieldName.split(";")[0];

      currentEvent[baseFieldName] = fieldValue;
      currentField = baseFieldName;
    }
  }

  return events;
}

function parseICSDateTime(icsDateTime) {
  // ICS format: 20260313T033000Z
  // Extract: YYYY-MM-DD HH:MM:SS
  const year = icsDateTime.substring(0, 4);
  const month = icsDateTime.substring(4, 6);
  const day = icsDateTime.substring(6, 8);
  const hour = icsDateTime.substring(9, 11);
  const minute = icsDateTime.substring(11, 13);
  const second = icsDateTime.substring(13, 15);

  return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
}

function escapeSQL(str) {
  if (!str) return "";
  return str.replace(/'/g, "''").replace(/\\/g, "\\\\");
}

function generateInsertStatements(events) {
  const statements = [];

  for (const event of events) {
    const summary = escapeSQL(event.SUMMARY || "");
    const location = escapeSQL(event.LOCATION || "");
    const startTime = event.DTSTART ? parseICSDateTime(event.DTSTART) : null;

    if (!summary || !location || !startTime) {
      console.warn("Skipping event with missing required fields:", event);
      continue;
    }

    const sql = `INSERT INTO \`session\` (\`summary\`, \`location\`, \`start_time\`, \`series\`) VALUES ('${summary}', '${location}', '${startTime}', 'f1');`;
    statements.push(sql);
  }

  return statements;
}

// Main execution
const args = process.argv.slice(2);

if (args.length === 0) {
  console.error("Usage: node parse-ics.js <path-to-ics-file> [output-file]");
  console.error("Example: node parse-ics.js calendar.ics output.sql");
  process.exit(1);
}

const inputFile = args[0];
const outputFile = args[1] || null;

try {
  const fileContent = readFileSync(inputFile, "utf8");
  const events = parseICS(fileContent);
  const insertStatements = generateInsertStatements(events);

  const output = insertStatements.join("\n");

  if (outputFile) {
    writeFileSync(outputFile, output, "utf8");
    console.log(
      `Generated ${insertStatements.length} INSERT statements in ${outputFile}`,
    );
  } else {
    console.log(output);
  }
} catch (error) {
  console.error("Error processing file:", error.message);
  process.exit(1);
}
