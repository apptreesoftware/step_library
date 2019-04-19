using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model {

    public class OtherCostType {
        public int Id { get; set; }
        public string Name { get; set; }
        public bool Active { get; set; }
        public bool TaxableFlag { get; set; }
        public string IncomeCategory { get; set; }
        public string ContraAccountNumber { get; set; }
        public string OtherCostGLAccount { get; set; }
        public int? GLAccountId { get; set; }
        public int? GLAccountIdForMarkup { get; set; }
        public int? StandardChargeAmount { get; set; }
        public bool DiscountedFlag { get; set; }
    }

}
