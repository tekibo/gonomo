import { formatName } from "../utils/helpers";

export default defineEventHandler(async (event) => {
  try {
    const name = (await readBody(event)).name as string;

    if (!name)
      throw createError({ statusCode: 400, message: "Name is required" });

    return { message: `Hello, ${formatName(name)}!` };
  } catch (error: any) {
    throw createError({ statusCode: 500, message: "Error: " + error.message });
  }
});
