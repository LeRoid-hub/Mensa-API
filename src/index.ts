// src/index.js
import express, { Express, Request, Response} from "express";
import dotenv from "dotenv";
import cors from "cors";
import bodyParser, { text } from "body-parser";
import morgan from "morgan";

import fetch from './fetch.js';
import Cache from './cache.js';
import {stripMensa, stripCampus} from "./stripper.js";

dotenv.config();

const app: Express = express().use(cors({ origin: '*' })).use(bodyParser.json());
app.use(morgan('combined'))

const port = process.env.PORT || 3000;

const baseUrl = "https://www.imensa.de/";

const cache = new Cache();

app.get("/", (req: Request, res: Response) => {
  res.send("Mensa API");
});


app.get("/api/bl/:Bundesland", (req: Request, res: Response) => {
    if (req.params.Bundesland === undefined) {
        return res.send("Invalid request");
    }
    let cachedData = cache.get("BL: "+req.params.Bundesland);
    if (cachedData !== null) {
        return res.send(cachedData);
    }else {
        let url = baseUrl + req.params.Bundesland.toLowerCase()+".html";
        fetch(url).then((data) => {
            let stripedData = null;
            if (data === null) {
                return res.send("Invalid request");
            }
            stripedData = stripCampus(data);
            cache.set("BL: "+req.params.Bundesland, stripedData);
            res.send(stripedData);
        });
    }
});
app.get("/api/:Location/:Mensa?", (req: Request, res: Response) => {
    if (req.params.Location === undefined) {
        return res.send("Invalid request");
    }
    let cachedData = cache.get(req.params.Mensa ?? req.params.Location);
    if (cachedData !== null) {
        return res.send(cachedData);
    }else {
        let url = baseUrl + req.params.Location.toLowerCase(); 
        if (req.params.Mensa !== undefined) {
            url += "/" + req.params.Mensa.toLowerCase();
        }
        fetch(url).then((data) => {
            let stripedData = null;
            if (data === null) {
                return res.send("Invalid request");
            }
            if (req.params.Mensa !== undefined) {
                stripedData = stripMensa(data);
            }else {
                stripedData = stripCampus(data);
            }

            cache.set(req.params.Mensa ?? req.params.Location, stripedData);
            res.send(stripedData);
        });
    }

});

    
app.get("/api", (req: Request, res: Response) => {
    res.send("/api/:Location/:Mensa?");
});

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});
