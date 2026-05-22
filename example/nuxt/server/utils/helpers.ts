export function formatName(name: string): string {
  return name
    .trim() // Remove leading/trailing spaces
    .replaceAll(/[!"#$%&'()*+,./:;<=>?@[\]^_`{|}~]/g, "") // Remove all punctuation
    .replace(/\s+/g, " ") // Collapse multiple spaces into one
    .split(" ") // Split into words
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase()) // Capitalize each word
    .join(" "); // Join back together
}
