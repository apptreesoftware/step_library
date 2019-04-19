using System;
using System.Collections.Generic;
using Newtonsoft.Json;

namespace Famis.Model
{
    public class InvoiceHeader
    {

        public int? Id { get; set; }

        public string ExternalId { get; set; }

        public string InvoiceNumber { get; set; }

        public string Description { get; set; }

        public DateTimeOffset InvoiceDate { get; set; }

        public int? VendorId { get; set; }

        public int? LienWaiverAmount { get; set; }

        public string LienWaiverComment { get; set; }

        public decimal? TaxAmount { get; set; }

        public decimal? ShippingAmount { get; set; }

        public decimal? TotalAmount { get; set; }

        public int? BudgetYear { get; set; }

        public DateTimeOffset DueDate { get; set; }

        public string VendorExternalId { get; set; }

        public List<InvoiceLine> Lines { get; set; }

        public string Message { get; set; }

        public string File { get; set; }

        [JsonProperty(PropertyName = "id")]
        public string IdAsString { get; set; }

        public bool ShouldSerializeLines()
        {
            return (false);
        }
        public bool ShouldSerializeMessage()
        {
            return (false);
        }
        public bool ShouldSerializeId()
        {
            return (false);
        }
        public bool ShouldSerializeIdAsString()
        {
            return (false);
        }
        public bool ShouldSerializeFile()
        {
            return (false);
        }
    }
}