import { type CollectionEntry } from "astro:content";

type Session = CollectionEntry<"sessions">["data"][number];

export function getLevelLabel(level: Session["level"]) {
  switch (level) {
    case "advanced":
      return "Advanced";
    case "all":
      return "All";
    case "beginner":
      return "Beginner";
    case "intermediate":
      return "Intermediate";
    default:
      throw new Error(`unknown session level: ${level}`);
  }
}

export function getTrackClass(track: Session["track"]) {
  switch (track) {
    case "room-1":
      return "room-1";
    case "room-2":
      return "room-2";
    default:
      throw new Error(`unknown session track: ${track}`);
  }
}

export function getTrackLabel(track: Session["track"]) {
  switch (track) {
    case "room-1":
      return "Room A";
    case "room-2":
      return "Room B";
    default:
      throw new Error(`unknown session track: ${track}`);
  }
}

export function getTypeLabel(type: Session["type"]) {
  switch (type) {
    case "lt":
      return "5 mins.";
    case "challenge":
    case "short":
      return "20 mins.";
    case "long":
      return "40 mins.";
    default:
      throw new Error(`unknown session type: ${type}`);
  }
}
