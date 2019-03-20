import axios from "axios"
import {getBaseURL} from "./constants";

const url = getBaseURL()

describe("Grabs info about the system", () => {
    it("Access base path", async () => {

        const resp = await axios.get(url)
        expect(resp.status).toEqual(200)
        expect(resp.data.status).toEqual("ok")

    })
})