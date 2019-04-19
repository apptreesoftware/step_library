using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model {
    public class OtherCost {
        public int OtherCostTypeId { get; set; }
        public Othercosttype OtherCostType { get; set; }
        public double Quantity { get; set; }
        public double UnitCost { get; set; }
        public double? ShippingAndHandlingCost { get; set; }
        public Account Account { get; set; }
        public int? ReceivedById { get; set; }
        public string ApInvoiceNumber { get; set; }
        public string ApAccountNumber { get; set; }
        public double? LienWaiverAmount { get; set; }
        public string LienWaiverComments { get; set; }
        public string PaymentComments { get; set; }
        public double? ApInvoiceAmount { get; set; }
        public DateTime? ApPostDate { get; set; }
        public DateTime? ApDueDate { get; set; }
        public double? ApTaxAmount { get; set; }
        public double? ApShippingHandlingAmount { get; set; }
        public int? ApInvoiceStatusId { get; set; }
        public DateTime? ApInvoiceDate { get; set; }
        public int? CoaAccountId { get; set; }
        public string Payee { get; set; }
        public double? TaxAmount { get; set; }
        public int? BudgetYear { get; set; }
        public string VendorCode { get; set; }
        public string VendorCompanyExternalId { get; set; }
        public int? Id { get; set; }
        public int? RequestId { get; set; }
        public string RequestExternalId { get; set; }
        public DateTimeOffset? Date { get; set; }
        public string Description { get; set; }
        public int? VendorCompanyId { get; set; }
        [JsonIgnore]
        public double TotalAmount { get; set; }
        public bool MarkupFlag { get; set; }
        public double? TotalMarkup { get; set; }
        public int? AccountId { get; set; }
        public DateTime? UpdateDate { get; set; }
        public int? RecurrenceId { get; set; }
        public int? UpdatedById { get; set; }
        //public bool AccountsPayableFlag { get; set; }
        //public bool ApExportFlag { get; set; }
        public bool TaxableFlag { get; set; }
        public string CurrencySign { get; set; }
        public string CurrencyCode { get; set; }
        public int? LineItemNumber { get; set; }
        public string ExternalId { get; set; }
        public double? TaxRate { get; set; }
        public string ExternalUpdatedById { get; set; }
        public string ExternalRecurrenceId { get; set; }
        public string InvoiceNumber { get; set; }
        public DateTime? InvoiceDate { get; set; }
    }

    public class Othercosttype {
        public int Id { get; set; }
        public string Name { get; set; }
        public bool Active { get; set; }
        public bool TaxableFlag { get; set; }
        public string IncomeCategory { get; set; }
        public string ContraAccountNumber { get; set; }
        public string OtherCostGLAccount { get; set; }
        public int GLAccountId { get; set; }
        public int GLAccountIdForMarkup { get; set; }
        public int StandardChargeAmount { get; set; }
        public bool DiscountedFlag { get; set; }
    }
}
