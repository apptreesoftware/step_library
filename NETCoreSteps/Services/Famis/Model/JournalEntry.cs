using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;

namespace Famis.Model
{
    public class JournalEntry
    {
        public int Id { get; set; }
        public string ExtractType { get; set; }
        public decimal Amount { get; set; }
        public int? DebitAccount { get; set; }
        public string DebitAccountDesc { get; set; }
        public int? CreditAccount { get; set; }
        public string CreditAccountDesc { get; set; }
        public DateTime? TransactionDate { get; set; }
        public DateTime? AccountingPeriod { get; set; }
        public DateTime? ExtractDate { get; set; }
        public DateTime? ExportDate { get; set; }
        public int? RequestId { get; set; }
        public int? TransactionId { get; set; }
        public string DebitEntity { get; set; }
        public string CreditEntity { get; set; }
        public string AccountErrors { get; set; }
        public string AdjustmentReason { get; set; }
        public int? DebitValuesAccount { get; set; }
        public int? CreditValuesAccount { get; set; }
        public string ChargeDescription { get; set; }
        public string SOW { get; set; }
        public int? DebitEntityAccountId { get; set; }
        public int? CreditEntityAccountId { get; set; }
        public bool EstimateFlag { get; set; }
        public string TransactionNumber { get; set; }
        public int? BatchId { get; set; }
        public int? WorkOrderId { get; set; }
        public string CostSource { get; set; }
    }
}