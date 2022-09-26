import { Request, Response, Router } from "express";

export const userRouter: Router = Router();

userRouter.get("/:id/get", async (req: Request, res: Response) => {
  return res.json("response");
});

userRouter.post("/create", async (req: Request, res: Response) => {
  return res.json("response");
});

userRouter.get("/get_all", async (req: Request, res: Response) => {
  return res.json("response");
});

userRouter.get(
  "/get_by_username/:username",
  async (req: Request, res: Response) => {
    return res.json("response");
  }
);

userRouter.delete("/drop", async (req: Request, res: Response) => {
  return res.json("response");
});
