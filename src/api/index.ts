import axios, { AxiosResponse } from "axios";

import { SpeedtestServer } from "./types";

// Search server via its name.
// Limit the search to the first 10 matching results.
export const searchByName = async (
	query: string
): Promise<AxiosResponse<SpeedtestServer[]>> =>
	axios.get(
		`https://www.speedtest.net/api/js/servers?search=${encodeURIComponent(
			query
		)}&limit=10`,
		{ timeout: 2000 }
	);
