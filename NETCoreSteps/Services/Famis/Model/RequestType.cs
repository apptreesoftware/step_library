using Newtonsoft.Json;

namespace Famis.Model
{
    public class RequestType
    {
        [JsonProperty("DefaultOpenStatusExternalId")]
        public object DefaultOpenStatusExternalId { get; set; }

        [JsonProperty("Id")]
        public long Id { get; set; }

        [JsonProperty("Description")]
        public string Description { get; set; }

        [JsonProperty("TabOrder")]
        public long TabOrder { get; set; }

        [JsonProperty("ActiveFlag")]
        public bool ActiveFlag { get; set; }

        [JsonProperty("ProjectFlag")]
        public bool ProjectFlag { get; set; }

        [JsonProperty("SurveyTypeId")]
        public object SurveyTypeId { get; set; }

        [JsonProperty("AssetReqCreateFlag")]
        public bool AssetReqCreateFlag { get; set; }

        [JsonProperty("ExternalId")]
        public string ExternalId { get; set; }

        [JsonProperty("AssetReqCloseFlag")]
        public bool AssetReqCloseFlag { get; set; }

        [JsonProperty("PushWoExtSystemFlag")]
        public bool PushWoExtSystemFlag { get; set; }

        [JsonProperty("DefaultOpenStatusId")]
        public object DefaultOpenStatusId { get; set; }

        [JsonProperty("MaintenanceFlag")]
        public bool MaintenanceFlag { get; set; }

        [JsonProperty("AllowParentChildFlag")]
        public bool AllowParentChildFlag { get; set; }

        [JsonProperty("StandingWoTypeFlag")]
        public bool StandingWoTypeFlag { get; set; }

        [JsonProperty("PsArInvoiceFlag")]
        public bool PsArInvoiceFlag { get; set; }

        [JsonProperty("SignatureReqWoCompleteFlag")]
        public bool SignatureReqWoCompleteFlag { get; set; }

        [JsonProperty("AccountStringRequiredCreateFlag")]
        public bool AccountStringRequiredCreateFlag { get; set; }

        [JsonProperty("AccountStringRequiredCloseFlag")]
        public bool AccountStringRequiredCloseFlag { get; set; }
    }
}