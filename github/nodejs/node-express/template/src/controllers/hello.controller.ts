import { Request, Response } from "express";

export function sayHello(req: Request, res: Response) {
  res.json({ message: "Hello, TypeScript!" });
};
