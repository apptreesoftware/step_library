using System;

namespace Famis.Model
{
    public class PrLine
    {
        
        public int Id { get; set; }
        
        public int? PRId { get; set; }
        
        public string PRNumber { get; set; }
        
        public int LineNumber { get; set; }
        
        public int? StatusId { get; set; }
        
        public DateTimeOffset RequiredDate { get; set; }
        
        public int? RequestId { get; set; }
        
        public string RequestExternalId { get; set; }
        
        public bool Material { get; set; }
        
        public bool OtherCostType { get; set; }
        
        public int? OtherCostTypeId { get; set; }
        
        public int? VendorId { get; set; }
        
        public string VendorExternalId { get; set; }
        
        public string VendorName { get; set; }
        
        public string VendorAddress { get; set; }
        
        public int? WarehouseId { get; set; }
        
        public string WarehouseExternalId { get; set; }
        
        public int? MaterialItemId { get; set; }
        
        public string MaterialItemExternalId { get; set; }
        
        public string PartNumber { get; set; }
        
        public int? BuyerId { get; set; }
        
        public string BuyerExternalId { get; set; }
        
        public string Description { get; set; }
        
        public int? UnitOfMeasureId { get; set; }
        
        public decimal? Quantity { get; set; }
        
        public decimal? UnitCost { get; set; }
        
        public DateTimeOffset CreateDate { get; set; }
        
        public string CreatedByName { get; set; }
        
        public DateTimeOffset UpdateDate { get; set; }
        
        public string UpdatedByName { get; set; }
        
        public int? GLAccountId { get; set; }
        
        public bool Active { get; set; }
       
        public string Message { get; set; }

        public string UomDescription { get; set; }

        public bool ShouldSerializeId()
        {
            return (false);
        }
        public bool ShouldSerializeMessage()
        {
            return (false);
        }
    }
}