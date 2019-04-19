using Newtonsoft.Json;
using System;

namespace Famis.Model
{
    public class PoLine
    {
        public int Id { get; set; }

        public int POLineNumber { get; set; }

        public string UpdatedById { get; set; }

        public string UpdatedByExternalId { get; set; }

        public DateTime? UpdateDate { get; set; }

        public string TotalAmount { get; set; }

        public int? POId { get; set; }

        public string PONumber { get; set; }

        public bool Active { get; set; }

        public decimal QuantityOrdered { get; set; }

        public string Description { get; set; }

        public DateTimeOffset? DateRequired { get; set; }

        public int? PropertyId { get; set; }

        public string PropertyExternalId { get; set; }

        public bool TaxableFlag { get; set; }

        public decimal TaxRate { get; set; }

        public decimal ShippingHandling { get; set; }

        public decimal UnitCost { get; set; }

        public bool OtherCostsFlag { get; set; }

        public int? OtherCostId { get; set; }

        public bool MaterialsFlag { get; set; }

        public int? UOMId { get; set; }

        public int? MaterialItemId { get; set; }

        public string MaterialItemExternalId { get; set; }

        public string MaterialItemPartNumber { get; set; }

        public int? WarehouseMaterialId { get; set; }

        public int? WarehouseId { get; set; }

        public string WarehouseExternalId { get; set; }

        public int? RequestId { get; set; }

        public string RequestExternalId { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public int? PRId { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public string PRNumber { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public int? PRLineId { get; set; }
        [JsonProperty(NullValueHandling = NullValueHandling.Ignore)]
        public int? PRLineNumber { get; set; }

        public bool ShouldSerializeId() {
            return false;
        }

        public bool ShouldSerializePRNumber() {
            return false;
        }

        public bool ShouldSerializePRId() {
            return MaterialsFlag;
        }

        public bool ShouldSerializePRLineId() {
            return MaterialsFlag;
        }

        public bool ShouldSerializePRLineNumber() {
            return false;
        }

        public bool ShouldSerializeUpdateDate() {
            return false;
        }

        public bool ShouldSerializeUpdatedById() {
            return false;
        }

        public bool ShouldSerializeUpdatedByExternalId() {
            return false;
        }

        public bool ShouldSerializeTotalAmount() {
            return false;
        }

        public bool ShouldSerializePOLineNumber() {
            return false;
        }

        public bool ShouldSerializeActive() {
            return false;
        }

        public bool ShouldSerializePropertyExternalId() {
            return false;
        }

        public bool ShouldSerializeRequestExternalId() {
            return false;
        }

        public bool ShouldSerializeWarehouseExternalId() {
            return false;
        }

        public bool ShouldSerializeMaterialItemExternalId() {
            return false;
        }
    }
}