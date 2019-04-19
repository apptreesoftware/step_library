using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model {
    public class CreateWorkOrder {
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public string ExternalId { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public int? RequestTypeId { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public int? RequestSubTypeId { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public int? SpaceId { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public string StatementOfWork { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public string ParentWOId { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public int? RequestPriorityId { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public int? CrewId { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public string FiscalYear { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public int? AssetId { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        [JsonIgnore]
        public Udf Udf { get; set; }
    }
}
