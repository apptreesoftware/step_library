using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model
{

    public class Warehouse
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public bool Active { get; set; }
        public int TabOrder { get; set; }
        public bool UseMovingCostAverage { get; set; }
        public bool StockItemIssueRequiredFlag { get; set; }
        public int? ExtractRegionId { get; set; }
    }

}
