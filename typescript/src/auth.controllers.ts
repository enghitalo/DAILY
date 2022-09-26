import { Request, Response, Router } from "express";
import { AuthService } from "./auth.services";

export const authRouter: Router = Router();

authRouter.post("/login", async (req: Request, res: Response) => {
  const authService: AuthService = new AuthService();
  const body = req.body;
  const response = authService.auth(body?.username, body?.password);
  return res.json("response");
});
