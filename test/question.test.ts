import axios from "axios";
import {getBaseURL} from "./constants";

const fs = require("fs").promises;

const relativePath = "/questions"
const url = getBaseURL() + relativePath

describe("Inserts and manipulates questions", () => {

    it("Inserts multiple questions",  async () => {

        const data = new Buffer(await fs.readFile("test/res/question.pdf"))
            .toString('base64')

        for(const i in ["a", "b", "c", "d"]) {
            const q = {
                "pdf": data,
                "answer": 0,
                "tags": [i, "ENEM"],
                "domain": i.toString()
            }
            const resp = await axios.post(url, q)
            expect(resp.status).toBe(201)
        }
        const resp = await axios.post(url, {
            "pdf": data,
            "answer": 1,
            "tags": ["0", "OTHER"],
            "domain": "0"
        })
        expect(resp.status).toBe(201)
    })

    it("Retrieve random questions filtering by one tag", async () => {
        const resp = await  axios.get(url + "?amount=5&tags[]=0")
        const data = resp.data
        expect(resp.status).toBe(200)
        expect(data.data.length).toBe(2)
        expect(data.data[1]["answer"] +
            data.data[0]["answer"]).toBe(1)
    })

    it("Retrieve random questions filtering by two tags", async () => {
        const resp = await  axios.get(url + "?amount=5&tags[]=0&tags[]=ENEM")
        const data = resp.data
        expect(resp.status).toBe(200)
        expect(data.data.length).toBe(1)
        expect(data.data[0]["answer"]).toBe(0)
    })

    it("Retrieve question using id", async () => {
        const initialResp = await  axios.get(url + "?amount=5&tags[]=0&tags[]=ENEM")
        const initialData = initialResp.data
        const id = initialData.data[0].id

        const resp = await axios.get(url + "/" + id)
        const data = resp.data
        expect(resp.status).toBe(200)
        expect(data.data.tags).toEqual(["0", "ENEM"])
    })

})
