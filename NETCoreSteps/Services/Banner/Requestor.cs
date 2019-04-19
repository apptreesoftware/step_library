using System;
using Newtonsoft.Json;

namespace Banner {
    public class Requestor
    {
        [JsonProperty("spriden_first_name")]
        public string FirstName { get; set; }

        [JsonProperty("spriden_mi")]
        public string MiddleInitial { get; set; }

        [JsonProperty("ftvorgn_title")]
        public string Department { get; set; }

        [JsonProperty("pebempl_orgn_code_home")]
        public string DepartmentCode { get; set; }

        [JsonProperty("spriden_id")]
        public string Id { get; set; }

        [JsonProperty("spriden_last_name")]
        public string LastName { get; set; }
    }
}