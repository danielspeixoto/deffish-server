import axios from "axios";
import {getBaseURL} from "./constants";

const fs = require("fs").promises;

const relativePath = "/question"
const url = getBaseURL() + relativePath

describe("Inserts and manipulates questions", () => {

    it("Inserts multiple questions",  async () => {

        const data = new Buffer(await fs.readFile("test/res/question.pdf"))
            .toString('base64')

        for(const i in ["a", "b", "c", "d"]) {
            const q = {
                "pdf": data,
                "answer": 0,
                "choices": ["A", "B", "C"],
                "tags": [i, "ENEM"]
            }
            const resp = await axios.post(url, q)
            expect(resp.status).toBe(201)
        }
        const resp = await axios.post(url, {
            "pdf": data,
            "answer": 1,
            "choices": ["A", "B", "C"],
            "tags": ["0"]
        })
        expect(resp.status).toBe(201)
    })

    it("Retrieve random questions filtering by one tag", async () => {
        const resp = await  axios.get(url + "?mode=random&amount=5&tags[]=0")
        const data = resp.data
        expect(resp.status).toBe(200)
        expect(data.data.length).toBe(2)
        expect(data.data[1]["answer"] +
            data.data[0]["answer"]).toBe(1)
    })
})
