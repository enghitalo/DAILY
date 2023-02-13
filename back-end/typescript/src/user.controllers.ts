import { Request, Response, Router } from "express";
import { User } from "./user.entities";
import { UserService } from "./user.services";

export const userRouter: Router = Router();

userRouter.get("/:id/get", async (req: Request, res: Response) => {
  const userService = new UserService();
  const user: User = userService.getUserByIdService(req.params.id);
  return res.json(user);
});

userRouter.post("/create", async (req: Request, res: Response) => {
  const userService = new UserService();
  return res.json("response");
});

userRouter.get("/get_all", async (req: Request, res: Response) => {
  const userService = new UserService();
  return res.json("response");
});

userRouter.get("/get_by_username/:username", async (req: Request, res: Response) => {
  const userService = new UserService();
  return res.json("response");
});

userRouter.delete("/drop", async (req: Request, res: Response) => {
  const userService = new UserService();
  return res.json("response");
});
