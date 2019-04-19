namespace Famis.Model
{
    public class InvoiceLine
    {
       
        public int Id { get; set; }
       
        public int InvoiceId { get; set; }
        
        public string InvoiceExternalId { get; set; }
       
        public int? WoDetailId { get; set; }
       
        public string WoDetailExternalId { get; set; }
        
        public int? PoDetailId { get; set; }
       
        public string PoDetailExternalId { get; set; }
       
        public int? LineItem { get; set; }
        
        public decimal Quantity { get; set; }
        
        public decimal Amount { get; set; }

        public string Description { get; set; }
        
        public decimal TaxAmount { get; set; }
        
        public decimal ShAmount { get; set; }

        public bool MatchStatus { get; set; }

        public int? MatchedReceiptWoId { get; set; }

        public string MatchedReceiptWoExternalId { get; set; }
       
        public string Message { get; set; }

        public bool ShouldSerializeId()
        {
            return (false);
        }
        public bool ShouldSerializeMessage()
        {
            return (false);
        }
        public bool ShouldSerializeLineItem()
        {
            return (false);
        }
        public bool ShouldSerializeDescription()
        {
            return (false);
        }
        public bool ShouldSerializeMatchStatus()
        {
            return (false);
        }
        public bool ShouldSerializePoDetailExternalId()
        {
            return (false);
        }
        public bool ShouldSerializeWoDetailId()
        {
            return (false);
        }
        public bool ShouldSerializeWoDetailExternalId()
        {
            return (false);
        }
    }
}