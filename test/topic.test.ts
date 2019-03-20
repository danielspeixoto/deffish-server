import axios from "axios";
import {getBaseURL} from "./constants";

const fs = require("fs").promises;

const relativePath = "/topic"
const url = getBaseURL() + relativePath

describe("Inserts and manipulates topics", () => {

    it("Inserts multiple topics",  async () => {

        for(const i in ["a", "b", "c", "d"]) {
            const q = {
                "title":  i,
                "contents": [
                    "A", i,
                ],
            }
            const resp = await axios.post(url, q)
            expect(resp.status).toBe(201)
        }
        const resp = await axios.post(url, {
            "title":  "e",
            "contents": [
                "A", "a",
            ],
        })
        expect(resp.status).toBe(201)
    })

    it("Retrieve random topics", async () => {
        const resp = await  axios.get(url + "?mode=random&amount=3")
        const data = resp.data
        expect(resp.status).toBe(200)
        expect(data.data.length).toBe(3)
    })

    it("Retrieve topic using id", async () => {
        const initialResp = await  axios.get(url + "?mode=random")
        const initialData = initialResp.data
        const id = initialData.data[0].id
        const title = initialData.data[0].title

        const resp = await axios.get(url + "/" + id)
        const data = resp.data
        expect(resp.status).toBe(200)
        expect(data.data.title).toEqual(title)
    })

})
